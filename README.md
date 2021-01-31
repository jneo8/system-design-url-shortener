# System Design URL Shortener 


## Requirements and Goals


### Functional requirements

* User give system the original url. System generate the unique and short url of it. The url must be short enough to be easily copy and pasted.
* When user access the short url. System should redirect them to the original url.
* Short url will expire after default time (2 years). User should able the choose the expiration time.
* User -> Login user or stranger. If same login user input same url -> return same short url.
* If multiple user or stanger input same url -> always return different url.
* Url performance visualization.

### Non-functional requirements

* Highly available
* Url redirection should happen in real-time with minimal latency.
* Shortened links should not be guessable (not predictable)
