# Api Race Challenge

In this challenge, you will need to apply what we've learned about `Multithreading` and `APIs` to `fetch the quickest result between two distinct APIs`.

Both requests will be made simultaneously to the following APIs:

1. <https://brasilapi.com.br/api/cep/v1/> + cep
1. <http://viacep.com.br/ws/> + cep + <"/json/>

The requirements for this challenge are:

1. Get the API that delivers the `fastest response` and discard the slower response.
1. The result of the request should be `displayed on the command line` with the address data, as well as which API sent it.
1. `Limit the response time to 1 second`. Otherwise, a timeout error should be displayed.
