/**
 * Go Video Downloader
 *
 *    Copyright 2018 Tenta, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * For any questions, please contact developer@tenta.io
 *
 * generator.go: Implementation of Python's generators
 *               using goroutines and channels
 */

package runtime

// Generator ..
type Generator interface {
	Object
	Yield(Object) Object
	Send(Object) Object
	run()
}

type generatorState int

const (
	generatorStateNew generatorState = iota
	generatorStateStarted
	generatorStateFinished
)

type generatorItem struct {
	val Object
	err interface{}
}

type generatorStruct struct {
	Object
	state generatorState
	fn    func(Generator) Object
	ch    chan generatorItem
}

func (g *generatorStruct) Yield(o Object) Object {
	g.ch <- generatorItem{val: o}
	in := <-g.ch
	if in.err != nil {
		panic(in.err)
	}
	return in.val
}

func (g *generatorStruct) Send(o Object) Object {
	switch g.state {
	case generatorStateNew:
		if o != None {
			panic(RaiseType(TypeErrorType,
				"can't send non-None value to a just-started generator"))
		}
		g.state = generatorStateStarted
		go g.run()
	case generatorStateFinished:
		panic(Raise(StopIterationType))
	}
	return g.Yield(o)
}

func (g *generatorStruct) run() {
	defer func() {
		if r := recover(); r != nil {
			g.state = generatorStateFinished
			g.ch <- generatorItem{err: r}
		}
	}()
	<-g.ch
	res := g.fn(g)
	g.state = generatorStateFinished
	g.ch <- generatorItem{err: Cal(StopIterationType, res)}
}

var _ Generator = (*generatorStruct)(nil)

// NewGenerator ..
func NewGenerator(fn func(Generator) Object) Generator {
	return &generatorStruct{
		Object: newObject(GeneratorType),
		state:  generatorStateNew,
		fn:     fn,
		ch:     make(chan generatorItem),
	}
}

// GeneratorType ..
var GeneratorType = newBuiltinType("generator")

func generatorIter(o Object) Object {
	return o
}

func generatorNext(o Object) Object {
	g := o.(Generator)
	return g.Send(None)
}

func generatorSend(args Args, kwArgs KWArgs) Object {
	checkFunctionArgs("send", args, kwArgs, GeneratorType, ObjectType)
	g := args[0].(Generator)
	o := args[1]
	return g.Send(o)
}

func initGeneratorType(slots *typeSlots, dict map[string]Object) {
	dict["send"] = newBuiltinFunction("send", generatorSend)
	slots.Iter = generatorIter
	slots.Next = generatorNext
}
