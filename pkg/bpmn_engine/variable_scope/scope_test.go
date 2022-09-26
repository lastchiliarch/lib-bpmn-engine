package variable_scope

import (
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
)

func Test_variableScope_GetVariable_Simple(t *testing.T) {
	// setup
	scope := NewScope(nil, nil)
	scope.SetVariable("name", "bpmn")

	// want
	then.AssertThat(t, scope.GetVariable("name"), is.EqualTo("bpmn"))
}

func Test_variableScope_GetVariable_Multiple(t *testing.T) {
	// setup
	rootScope := NewScope(nil, nil)
	rootScope.SetVariable("a", 1)
	rootScope.SetVariable("b", 2)

	taskCScope := NewScope(rootScope, nil)
	taskCScope.SetVariable("d", 4)

	taskSubProcess := NewScope(rootScope, nil)
	taskSubProcess.SetVariable("c", 3)
	taskAScope := NewScope(taskSubProcess, nil)
	taskAScope.SetVariable("b", 4)
	taskBScope := NewScope(taskSubProcess, nil)

	// want
	// check if root can see variable defined in root scope
	then.AssertThat(t, rootScope.GetVariable("a"), is.EqualTo(1))
	// check if b is overwritten by taskA
	then.AssertThat(t, rootScope.GetVariable("b"), is.EqualTo(2))

	// d should be visible to taskC
	then.AssertThat(t, taskCScope.GetVariable("d"), is.EqualTo(4))
	// d defined in taskC  should not visible to taskA
	then.AssertThat(t, taskAScope.GetVariable("d"), is.EqualTo(nil))

	// a defined in root should visible to taskC
	then.AssertThat(t, taskCScope.GetVariable("a"), is.EqualTo(1))

	// check variable defined in subProcess
	then.AssertThat(t, taskSubProcess.GetVariable("c"), is.EqualTo(3))
	then.AssertThat(t, taskSubProcess.GetVariable("b"), is.EqualTo(2))

	// b defined in both taskA and root scope, should overwrite by value in taskA
	then.AssertThat(t, taskAScope.GetVariable("b"), is.EqualTo(4))
	// b defined in root should visible to taskB
	then.AssertThat(t, taskBScope.GetVariable("b"), is.EqualTo(2))

}

func Test_variableScope_GetContext(t *testing.T) {
	// setup
	scope := NewScope(nil, nil)
	scope.SetVariable("name", "bpmn")

	// want
	then.AssertThat(t, scope.GetContext(), is.EqualTo(map[string]interface{}{"name": "bpmn"}))
}

func TestMergeScope(t *testing.T) {
	// setup
	rootScope := NewScope(nil, map[string]interface{}{
		"a": 1,
		"b": 2,
	})
	local := NewScope(nil, map[string]interface{}{
		"b": 3,
	})

	newScope := MergeScope(local, rootScope)

	// want
	then.AssertThat(t, newScope.GetContext(), is.EqualTo(map[string]interface{}{
		"a": 1,
		"b": 3,
	}))
}