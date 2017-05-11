Changes
=================================================

[Fix] - bugfix

**[Breaking]** - breaking change

[Feature] - new feature

[Improvement] - non-breaking improvement

[Code] - code quality related change that shouldn't make any significant difference for end-user


Changes
-------
**Master**
   - **[Breaking]** Logging migrated to zap (structured logging). Log format changed significantly. Old command line options removed. Please consult example.conf for a new config options and explanations
   - [Improvement] Add context support. Also log context from carbonapi
   - [Improvement] Add memcache support for zipper caches
   - [Improvement] Use dep as a vendoring tool
   - [Improvement] Add a Makefile that will hide some magic from user
   - [Improvement] Split carbonzipper into several packages
   - [Fix] Fix incompatibility between carbonzipper and older versions of carbonserver/go-carbon (protobuf2-only)

**0.63**
   - [Fix] carbonsearch query cache was never cleared

**0.62**
   - [Fix] Fix carbonsearch queries with recent carbonapi version
   - [Fix] Fix pathCache to handle render requests with globs.
   - [Feature] Add cache for carbonsearch results

**0.61**
   - [Fix] Fix rewrite for internal queries, because of an error some queries were sent as protobuf not as protobuf3
   - [Code] gofmt the code!

**0.60**
   - **[Breaking]** Carbonzipper backend protocol changed to protobuf3. Though output for /render, /info /find can be both (format=protobuf3 for protobuf3, format=protobuf for protobuf2).

**0.50**
   - See commit log.
