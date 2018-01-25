# Plan to get rid of Cobra

* Write an `input` package around `viper`. Use `input` to retrieve data:
  * from files
  * from environment
  * from arguments
  * from flags

* Write a `cli` package for all code parsing data from `input` and calling actual application code.