## Bakingdish

BakingDish provides a web interface for storing, searching and viewing
recipes.  Bakingdish is an exploratory project which exists so that I
can test out a number of interesting technologies together.  They
include

    * go itself, and the gocraft framework
    * angularjs and css 
    * OAuth2 and other authentication frameworks
    * mongodb
    * docker and its networking and composition tools
    * digital ocean, aws, gce and other providers

The various components are an angularjs client, a gocraft web api, a
nosql database backend store, with the front- and back-end components
hosted in docker containers, wired together with swarm or kubernetes
or other mechanisms.


### Go

After a little bit of time working with
[Beego](https://github.com/astaxie/beego) I've decided to go with a
lighter-weight framework, [gocraft](https://github.com/gocraft/web).
As a go learner, I'm better off with a codebase that diverges less
from the fundamentals of the language.  As well, these open-source
projects often have some sharp edges that require delving in to the
source code and documentation.  I found that the language barrier for
Beego was an additional hindrance that I don't need.

Gocraft is pretty straightforward about its usage of middleware and
routers, and so far has been approachable

### AngularJS

Well. Angular and client-side js will be good practice.

### Authentication and Authorization

t.b.d.

### Data storage

Initial tests with the mgo driver show it to be very usable.  Other
backends can be swapped in or out for testing.

### Containers

Docker is a fast-moving system, and I'd like to stay on top of the
developments in orchestration and networking.

There are interesting tricks, also, in building minimally-sized
containers.  In my build process I plan to construct a container for
the gocraft server binary which is as small as possible.  Ideally the
go binary will be statically compiled but it remains to be seen whether
that's doable with a dependence on net/http.

I don't expect that this service will have more than a few users, but
nevertheless it will be interesting to experiment with dynamically
scaling the nodes.