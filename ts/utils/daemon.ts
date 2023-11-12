import {
    Interceptor,
    PromiseClient,
    StreamRequest,
    StreamResponse,
    Transport,
    UnaryRequest,
    UnaryResponse,
    createPromiseClient,
} from '@connectrpc/connect';
import { AppDaemon } from '../v1/app_connect';

/**
 * NamespaceHeader designates the header used to specify the namespace.
 */
const NamespaceHeader = 'x-webmesh-namespace';
/**
 * DefaultNamespace is the default namespace.
 */
const DefaultNamespace = 'webmesh';
/** 
 * DefaultDaemonAddress is the default daemon address.
 */
const DefaultDaemonAddress = 'http://127.0.0.1:58080';

/**
 * DaemonClient is a type alias for the PromiseClient of the AppDaemon service.
 */
export type DaemonClient = PromiseClient<typeof AppDaemon>;

/**
 * DaemonOptions are the options for communicating with the daemon.
 */
export interface DaemonOptions {
    daemonAddress: string;
    namespace: string;
    transport: Transport;
    pollInterval?: number;
}

/**
 * Options are the options for using Webmesh. They are a superset of
 * DaemonOptions and can be used to inherit the daemon address and
 * namespace from the environment. It should be extended to provide
 * the transport needed to communicate with the daemon.
 */
export class Options implements DaemonOptions {
    daemonAddress: string;
    namespace: string;
    transport: Transport;
    pollInterval: number;

    static default(): Options {
        return new Options();
    }

    constructor(opts?: Partial<DaemonOptions>) {
        if (opts?.pollInterval && opts.pollInterval > 0) {
            this.pollInterval = opts.pollInterval;
        } else {
            this.pollInterval = 5000;
        }
        if (!opts?.transport) {
            throw new Error('transport is required');
        }
        this.transport = opts.transport;
        this.daemonAddress = opts?.daemonAddress ?? DefaultDaemonAddress;
        this.namespace =
            opts?.namespace ?? process.env.npm_package_name ?? DefaultNamespace;
    }

    /**
     * Interceptor returns an interceptor that sets the namespace header on RPC requests.
     * @returns the interceptor for the daemon.
     */
    public interceptor(): Interceptor {
        return (next) =>
            async (
                req: UnaryRequest | StreamRequest,
            ): Promise<UnaryResponse | StreamResponse> => {
                req.header.set(NamespaceHeader, this.namespace);
                return next(req);
            };
    }

    /**
     * client returns a client for the daemon.
     * @returns the client for the daemon.
     */
    public client(): DaemonClient {
        return createPromiseClient(AppDaemon, this.transport);
    }
}
