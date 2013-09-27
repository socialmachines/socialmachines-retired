Memory Model
============

Heap should only be accessible via channel.  HeapInsert and HeapLookup should be functions that insert or retrieve to/from the heap and block until the value is inserted or retrieved.

Scopes should be passed in by value and not by reference. After this then can then be chained as opposed to copied as they are now. Self needs to be added to scope to fix the current problems caused by passing it along in a message as the 0th argument.

The objects (and promises espcially) need to implement a reference counting scheme so that the goroutines (which currently run indefinately) will kill themselves when no longer needed (i.e. out of scope).

AST / Runtime
=============

The Expr interface methods (Eval, and Visit) should be changed to return uint64.  Currently, by returning Value, they datastructures loose information as was discovered by the implementation of Integer.


soma eval and :load
===================
```soma eval``` and ```:load``` need to be completed.  ```soma eval``` currently has a ```soma console``` option and just requires a completion of the implemention.  ```:load``` needs to be added to the console. ```:load``` will take one argument, the path to a file, and load it into the REPL.

define
======
Removed the block braces from definitions. Statements will be comma-seperated, and end with a period.

```
+ Foo bar: aBaz => "Foo is not accessible".

"Doing this will no longer make 'self' a special variable"
+ (f Foo) bar: aBaz => f doSomething.


+ (f Foo) bar: aBaz =>
  "Attributes are just a special type of unary or keyword message."

  "set the attribute @attribute on f with a value of aBaz"
  f @attribute: aBaz,

  "assign q to the value of @attr2 from object f"
  q := f @attr2.
```