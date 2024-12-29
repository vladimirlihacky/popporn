import * as config from "./config.json"
import { join } from "node:path"
import { pathToRegexp } from "path-to-regexp"
import { mkdir, readdir } from "node:fs/promises";

type ServiceParams = {
    root: string
}

class Service {

    root: string

    constructor({
        root
    }: ServiceParams) {
        this.root = root;

        this.checkRoot();
    }

    private async checkRoot() {
        try {
            await readdir(this.root)
        } 
        catch(err) {
            await mkdir(this.root)
        }
    }

    private createUUID() {
        return crypto.randomUUID()
    }

    public async handleUploadFile(req: Request): Promise<Response> {
        const formData = await req.formData()
        const file = formData.get("file")

        if (!(file instanceof File)) {
            return new Response(JSON.stringify({
                status: 'Error',
                message: 'No file uploaded'
            }), { status: 400 })
        }

        const fileName = this.createUUID()
        const path = join(this.root, fileName)

        await Bun.write(path, file)

        return new Response(JSON.stringify({
            status: 'OK',
            filename: fileName,
        }))
    }

    public async handleGetFile(uuid: string): Promise<Response> {
        const path = join(this.root, uuid);

        const file = await Bun.file(path)

        if (!await file.exists()) {
            return new Response(JSON.stringify({
                status: "Error",
                message: "File not found"
            }), { status: 404 })
        }

        return new Response(file)
    }

    public async handleDefaultRequest(req: Request): Promise<Response> {
        return new Response(JSON.stringify({
            status: "Error",
            message: "Not found."
        }), { status: 404 })
    }
}

const service = new Service({
    root: config.root
})

Bun.serve({
    port: config.port,
    hostname: config.host,

    async fetch(req) {
        const url = new URL(req.url);

        if (url.pathname === "/upload" && req.method === "POST") {
            return service.handleUploadFile(req)
        }

        const getFilePattern = pathToRegexp("/files/:uuid")
        const matches = getFilePattern.regexp.exec(url.pathname)
        if (matches && req.method === "GET") {
            const uuid = matches[1]

            return service.handleGetFile(uuid);
        }


        return service.handleDefaultRequest(req);
    }
});