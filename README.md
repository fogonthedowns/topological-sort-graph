# Topological Sort

We will use a Directed Graph.

```
directed graph:
a -> v
a -> c
b -> d
b -> c
d -> z
z -> f
c -> e
c -> q
q -> r
e -> f
```

In order to track the sorted list we will use a linked list.
One could think of the topological sort as a dependency build tree. for instance "c" depends on "a" and "e" depends on "c".

to track all the visited nodes This implementation uses a map of string -> bool.
