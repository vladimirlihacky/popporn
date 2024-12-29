import routes from "./routes.json"
import config from "./config.json"

type Route = {
    pathPrefix: string, 
    redirectTo: string, 
    allowedMethods: string[],
}

type ServiceArgs = {
    port: number,
    hostname: string,
}

class Service {

    private port: number; 
    private hostname: string;
    private routes: Route[] = [];

    constructor({
        port, 
        hostname,
    }: ServiceArgs) {
        this.port = port;
        this.hostname = hostname;
    }

    public addRoute(route: Route) {
        this.routes.push(route);
    }

    private matchRoute(url: string): Route | null {
        let matchedRoute: Route | null = null; 

        this.routes.forEach(route => {
            if(url.startsWith(route.pathPrefix)) {
                matchedRoute = route
            }
        })

        return matchedRoute
    }

    public start() {
        const service = this;

        Bun.serve({
            port: this.port, 
            hostname: this.hostname, 

            async fetch(req) {
                const url = req.url;
                const route = service.matchRoute(url)

                if(!route) {
                    return new Response(JSON.stringify({
                        status: "Error",
                        message: "Route not found"
                    }))
                }

                const path = url.slice(route.pathPrefix.length);

                return new Response(null, {
                    status: 302,
                    headers: {
                        location: `${route.redirectTo}/${path}`
                    }
                });
            }
        })
    }
}

const service = new Service({
    port: config.port,
    hostname: config.hostname,
})

routes.forEach(service.addRoute)

service.start()