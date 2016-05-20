# Creating a messaging server

According to the WebRTC specification, peers need to exchange messages before being able to connect to each other. However, the implementation of the signalling exchange is unspecified[1].


The propose of the exercise is to implement such a service:
* if A sends a message M to B and B is known by the service, the service should deliver M
to B
* the client should be implementable in a web browser
*  The implementation technology is open: any language or middleware (databases,
messaging system...) can be used. However, if a third-party system is used, the implementation must include a standalone version.

Once the implementation is done, think about the following questions:
* Which share of web users will this implementation address? Should it be increased and
how?
* How many users can connect to one server?
* How can the system support more users?

## Deliverables:

* Use a source management system (git and github prefered, but any other alternative is
accepted) and give us access to it.
* Please do not use Streamroot inside your project as is it a registred trademark.
* The submission must include code source AND build scripts if applicable.
* The submission may include a beginning of an answer to the follow-up questions above.
In any case, these questions will be discussed in the interview, so we strongly advise you give them some thought.

## Evaluation:

* Correctness of the solution
* Quality of the code
* Development practices (usage of the scm, tests...)
 