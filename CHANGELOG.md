# Changelog

## [1.1.0](https://github.com/Ay-afk-stack/gopher-socials/compare/v1.0.0...v1.1.0) (2026-05-14)


### Features

* update api version automatically ([d00b93c](https://github.com/Ay-afk-stack/gopher-socials/commit/d00b93c7b6fc516198417c5fd208b68f467706b0))

## 1.0.0 (2026-05-14)


### Features

* add automation workflow ([f02ce91](https://github.com/Ay-afk-stack/gopher-socials/commit/f02ce9107bd7381edc537fd2b42ee587f9bee93a))
* added a structured logger using zap ([5ea20b3](https://github.com/Ay-afk-stack/gopher-socials/commit/5ea20b3c5ee4ec1c1ae507840e4124804d09c0d1))
* added basic auth middleware ([5c0d0c9](https://github.com/Ay-afk-stack/gopher-socials/commit/5c0d0c9c21d052a64022bccd68cc8794810d4d25))
* added hashing for pw and token for invitation ([f6c72df](https://github.com/Ay-afk-stack/gopher-socials/commit/f6c72df291998010d1c96808622e8f017c3f3a50))
* added migrations of the roles and role_id in users ([7231c16](https://github.com/Ay-afk-stack/gopher-socials/commit/7231c16af1ca5a1f0078a7853ff1315970b34c9b))
* cache user data with Redis ([568dd4a](https://github.com/Ay-afk-stack/gopher-socials/commit/568dd4ad3852ba84c17a544a75fea1ab9916a7c6))
* changed the datatype of password and implemented users_invitation migration ([cc1f598](https://github.com/Ay-afk-stack/gopher-socials/commit/cc1f5989fd9921649d6dda3b191aca8fb47b76cb))
* completed follow and unfollow features ([434392b](https://github.com/Ay-afk-stack/gopher-socials/commit/434392b9a5aa1455e34017689f31cdd894afce90))
* configured env variables ([1850a87](https://github.com/Ay-afk-stack/gopher-socials/commit/1850a877786578a1b5ff410cce10396553f36dca))
* create users and post function implemented ([87678fb](https://github.com/Ay-afk-stack/gopher-socials/commit/87678fb3d9fde9d76a7d49a6d6acd2e8d41b1221))
* created connection for database using pgxpool ([f0eda7c](https://github.com/Ay-afk-stack/gopher-socials/commit/f0eda7cd1d888910d175ae27081af16ca498fbb7))
* created indexes for database ([5750c74](https://github.com/Ay-afk-stack/gopher-socials/commit/5750c74c068004aaee10df4cbe7536d047a13589))
* database migrated successfully! ([a1479ed](https://github.com/Ay-afk-stack/gopher-socials/commit/a1479ed7750f3024702a24546701a16d381b95ab))
* email verification implemented ([1c57c42](https://github.com/Ay-afk-stack/gopher-socials/commit/1c57c428ccd8e9bccbbcb98a6afa3b9e7ff1cc58))
* get posts with comments feature implemented ([c92390f](https://github.com/Ay-afk-stack/gopher-socials/commit/c92390f4b1948ff36ba0566eaa910e9e317427c3))
* get user by id implemented ([55d174a](https://github.com/Ay-afk-stack/gopher-socials/commit/55d174afc433f2185ed033746f81b2770f8e78bb))
* implemented activate functionality for users using token ([29687dd](https://github.com/Ay-afk-stack/gopher-socials/commit/29687dd00fd5a9827e80487ed1fe38c2f396ea01))
* implemented create and getbyid functionality for post ([fc520cd](https://github.com/Ay-afk-stack/gopher-socials/commit/fc520cd0c03b74e1b0d24f63990b2b170fdd8a38))
* implemented feed filtering using tags, search ([2e4aff3](https://github.com/Ay-afk-stack/gopher-socials/commit/2e4aff32c213d1126d7608099167691a0e544772))
* implemented feeds algorithm for users ([61b4c2f](https://github.com/Ay-afk-stack/gopher-socials/commit/61b4c2faa516a8a836b401cec9e9f79f9ccd4d5b))
* implemented graceful shutdown ([c0c8443](https://github.com/Ay-afk-stack/gopher-socials/commit/c0c84438b3a3481e2051deb2584c187f99731c06))
* implemented JWT authentication ([5ea10e6](https://github.com/Ay-afk-stack/gopher-socials/commit/5ea10e641d18cd6da9cdae0cdb041d583e351806))
* implemented pagination and sorting ([fd65d87](https://github.com/Ay-afk-stack/gopher-socials/commit/fd65d879e39edbba0a628b38e3149754705ae405))
* implemented role based authentication ([8c4d91d](https://github.com/Ay-afk-stack/gopher-socials/commit/8c4d91d953b18453245a8176ee9f0e8b781fe467))
* implemented update and delete with middleware ([6f15743](https://github.com/Ay-afk-stack/gopher-socials/commit/6f15743b00a6747abd8528c3da89b9a983a0c7f6))
* implemented validator package to validate payloads ([d1b2b1f](https://github.com/Ay-afk-stack/gopher-socials/commit/d1b2b1ffe96549c2a7717b2b039db6230abd3ad2))
* integrate chi router and add air for live reload ([0120a2b](https://github.com/Ay-afk-stack/gopher-socials/commit/0120a2ba2c1511dee3766a7a5dfcd6f0714433ce))
* release please scripts ([3ad6282](https://github.com/Ay-afk-stack/gopher-socials/commit/3ad62827a5c6ef39f5ab0a9bf7b657353fb1e6a2))
* server init ([2759622](https://github.com/Ay-afk-stack/gopher-socials/commit/27596225c282ca1d382f8d9be3aa443733a70139))
* verfied JWT and implemented auth middleware ([6a96cf9](https://github.com/Ay-afk-stack/gopher-socials/commit/6a96cf9b08c869aca170899feb29fe05406ca1db))


### Bug Fixes

* added password comparision in createtokenhandler ([8507676](https://github.com/Ay-afk-stack/gopher-socials/commit/85076769d8750467d3cf50edde173b992670c842))
* added staticchecks ([e296e48](https://github.com/Ay-afk-stack/gopher-socials/commit/e296e489b758680a8eb658ff7bbcc1e9252595b1))
* fixed concurrency error in update function ([ded9c6e](https://github.com/Ay-afk-stack/gopher-socials/commit/ded9c6e9a8df8b54c442dc1d05535efea180be02))
* fixed posts table error for migration ([3b7d420](https://github.com/Ay-afk-stack/gopher-socials/commit/3b7d420ea73580879617ce8726fd842fe402fff7))
* fixed the manual role id issue ([75c1f54](https://github.com/Ay-afk-stack/gopher-socials/commit/75c1f54f975645ff087906f8e7697a223e13465e))
* fixed the token string in users store ([6edd001](https://github.com/Ay-afk-stack/gopher-socials/commit/6edd0012d4a83f56171483d2338514c97636c21f))
* removed punctuation and unused variable ([9c8f6b7](https://github.com/Ay-afk-stack/gopher-socials/commit/9c8f6b7bab374f52b394deea2f6b41a58c0f0239))
* removed user context for staticcheck ([393e1f1](https://github.com/Ay-afk-stack/gopher-socials/commit/393e1f1781365ff243653868a44f168a5b6af0a4))
