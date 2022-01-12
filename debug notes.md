
# Notes for vecty and vecty-material

## did the 'child Render' illegally return a stored render variable?

### list.go

A crash is caused by `func ItemDivider() vecty.ComponentOrHTML`
demos\menu issues #33

### dom.go

line 509

```go
// Determine the next child render.
nextChildRender, skip, mounters := render(nextChild, prevChild)
if nextChildRender != nil && prevChildRender != nil && nextChildRender == prevChildRender {
    log.Printf("%v", "vecty/dom/vecty.HTML.reconcileChildren()1 vecty: next 'child render' must not equal previous 'child render' (did the 'child Render' illegally return a stored render variable?)")
    log.Printf("%v %v %p %+v %v %p %+v %v %v %v %v", "vecty/dom/vecty.HTML.reconcileChildren()1", "nextChild =", nextChild, nextChild, "prevChild =", prevChild, prevChild, "nextChildRender =", nextChildRender, "prevChildRender =", prevChildRender)
    //panic("vecty: next 'child render' must not equal previous 'child render' (did the 'child Render' illegally return a stored render variable?)")
}
```

line 646

```go
func (h *HTML) removeChild(child *HTML) {
// If we're removing the current insert target, use the next
// sibling, if any.
log.Printf("%v %v %p %+v %v %+v %v %p %+v", "vecty/dom/vecty.HTML.removeChild()1", "child =", child, child, "child.node =", child.node, "h =", h, h)
if h.insertBeforeNode != nil && h.insertBeforeNode.Equal(child.node) {
    h.insertBeforeNode = h.insertBeforeNode.Get("nextSibling")
}
unmount(child)
if child.node == nil {
    return
}
// Use the child's parent node here, in case our node is not a valid
// target by the time we're called.
pn := child.node.Get("parentNode")
log.Printf("%v %v %+v", "vecty/dom/vecty.HTML.removeChild()2", "child.node.Get('parentNode') =", pn)
child.node.Get("parentNode").Call("removeChild", child.node)
}
```

### Related issues

Panic when component uses ComonentOrHTML field as child #191
<https://github.com/hexops/vecty/issues/191>

Issues related to the error `next 'child render' must not equal previous 'child render'`
Design: Determine how to manage component lifecycle: <https://github.com/hexops/vecty/issues/115>

## mcdtest

This was removed after the following commit:
<https://github.com/vecty-components/material/tree/f179a1bb3ed41eba7a776f441d556289bedf43e1/components/internal/mdctest>

It seems to be required for testing.
