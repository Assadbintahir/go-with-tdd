### Go With TDD
This codebase is to track performance for doing
 [TDD with Golang](https://quii.gitbook.io/learn-go-with-tests/). This is an effort more close to applying TDD Paradigm in golang in comparison to implementing some feature set in golang.

DISCLAIMER: This repo may contain testing techniques outside of the scope of the attached guide.

##### Techniques

- Unit tests i.e. `t *testing.T`
- Unit tests with table test format
- Benchmarking of iterations i.e. `b *testing.B`
- Example tests showing example and asserting returned output
- Coverage calculation and enforcement i.e. `m *testing.M` (Helper function which enforces coverage threshold is extracted in a common util function. It can be reused in any package with appropriate parameters.)