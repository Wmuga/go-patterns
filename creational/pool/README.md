# Object Pool Pattern

The object pool creational design pattern is used to prepare and keep multiple
instances according to the demand expectation.


## Rules of Thumb

- Object pool pattern is useful in cases where object initialization is more
  expensive than the object maintenance.
- If there are spikes in demand as opposed to a steady demand, the maintenance
  overhead might overweigh the benefits of an object pool.
- It has positive effects on performance due to objects being initialized beforehand.
