## tspanner

Bigtable-like versioned key-value store

Data is stored in schematized semirelational tables



**features**

* provides externally consistent read and writes and globally consistent read  across the database at a timestamp



externally consistent ？（equivalently, linearizability [Herlihy and Wing 1990])）



group内通过2PL和paxos来实现transaction

不同的group通过2PC来实现



​	