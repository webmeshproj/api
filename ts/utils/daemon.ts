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
export const NamespaceHeader = 'x-webmesh-namespace';
/**
 * DefaultNamespace is the default namespace.
 */
export const DefaultNamespace = 'global';
/** 
 * DefaultDaemonAddress is the default daemon address.
 */
export const DefaultDaemonAddress = 'http://127.0.0.1:58080';

/**
 * DaemonClient is a type alias for the PromiseClient of the AppDaemon service.
 */
export type DaemonClient = PromiseClient<typeof AppDaemon>;

/**
 * DaemonOptions are the options for communicating with the daemon.
 */
export interface DaemonOptions {
    /**
     * daemonAddress is the address of the daemon.
     */
    daemonAddress: string;
    /**
     * namespace is the namespace to use for the daemon.
     */
    namespace: string;
    /**
     * transport is the transport to use for communicating with the daemon.
     */
    transport: Transport;
    /**
     * pollInterval is the interval in milliseconds to poll the daemon for updates.
     */
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

    /**
     * default returns the default options.
     * @returns the default options.
     */
    static default(): Options {
        return new Options();
    }

    /**
     * interceptor returns an interceptor that sets the namespace header on RPC requests.
     * @returns the interceptor for the daemon.
     */
    static interceptor(namespace: string): Interceptor {
        return (next) =>
            async (
                req: UnaryRequest | StreamRequest,
            ): Promise<UnaryResponse | StreamResponse> => {
                req.header.set(NamespaceHeader, namespace);
                return next(req);
            };
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
     * client returns a client for the daemon.
     * @returns the client for the daemon.
     */
    public client(): DaemonClient {
        return createPromiseClient(AppDaemon, this.transport);
    }
}
