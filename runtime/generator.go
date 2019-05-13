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

import (
	"errors"
	"math/rand"
	"runtime"
	"sync"
)

// useFinalizer governs wether abandoned generators should be
// stopped using their finalizers.
const useFinalizer = false

// Generator ..
type Generator interface {
	Object
	Send(Object) Object
}

// Yielder ..
type Yielder interface {
	Yield(Object) Object
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

type generatorCore struct {
	id    int
	state generatorState
	ch    chan generatorItem
	fn    func(Yielder) Object
}

type generatorStruct struct {
	Object
	*generatorCore
}

var errGeneratorFinalized = errors.New("Generator finalized")

func (g *generatorCore) Yield(o Object) Object {
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

func (g *generatorCore) run() {
	generatorStore.add(g)
	defer generatorStore.remove(g)
	defer func() {
		if r := recover(); r != nil {
			if r == errGeneratorFinalized {
				// do nothing, just finish run() asap
				return
			}
			g.state = generatorStateFinished
			g.ch <- generatorItem{err: r}
		}
	}()
	<-g.ch
	res := g.fn(g)
	g.state = generatorStateFinished
	g.ch <- generatorItem{err: Cal(StopIterationType, res)}
}

func (g *generatorCore) finalize() {
	if g.state == generatorStateStarted {
		// Inject a predefined error, which causes a stuck Yield running
		// in the associated goroutine to blow up, taking the entire
		// routine with it and finally terminating it cleanly in its
		// panic handler.
		g.ch <- generatorItem{err: errGeneratorFinalized}
	}
}

var _ Generator = (*generatorStruct)(nil)

// NewGenerator ..
func NewGenerator(fn func(Yielder) Object) Generator {
	g := &generatorStruct{
		Object: newObject(GeneratorType),
		generatorCore: &generatorCore{
			state: generatorStateNew,
			ch:    make(chan generatorItem),
			fn:    fn,
		},
	}
	// This finalizer is meant to clean up the run() routine when the
	// generator gets started but is not exhausted.
	// Note: we detached the innards of the generator into *generatorCore,
	// so that the generator can be garbage-collected even when it's running,
	// and the associated goroutine still references its core.
	if useFinalizer {
		runtime.SetFinalizer(g, func(g *generatorStruct) {
			g.finalize()
		})
	}
	return g
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

// ---- Generator store ----

var generatorStore = generatorStoreImpl{m: make(map[int]*generatorCore)}

type generatorStoreImpl struct {
	sync.Mutex
	m map[int]*generatorCore
}

func (gs *generatorStoreImpl) add(g *generatorCore) {
	gs.Lock()
	defer gs.Unlock()
	for {
		g.id = rand.Int()
		_, ok := gs.m[g.id]
		if !ok {
			break
		}
	}
	gs.m[g.id] = g
}

func (gs *generatorStoreImpl) remove(g *generatorCore) {
	gs.Lock()
	defer gs.Unlock()
	delete(gs.m, g.id)
}

// FinalizeGenerators kills all running generators.
func FinalizeGenerators() {
	generatorStore.Lock()
	defer generatorStore.Unlock()
	for _, g := range generatorStore.m {
		g.finalize()
	}
}
