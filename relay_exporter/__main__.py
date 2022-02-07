"""Simple Prometheus relay exporter entry-point module."""

import asyncio
import logging
from argparse import ArgumentParser, ArgumentDefaultsHelpFormatter
import uvloop
from yarl import URL
from aiohttp import web
from .handlers import relay_handler
from .signals import client_session_cleanup_ctx
from .version import __version__


DEF_HOST = "0.0.0.0"
DEF_PORT = 9878
DEF_CLIENT_PROXY = None
DEF_CLIENT_READ_SIZE = 2048
DEF_LOG_LEVEL = "INFO"
DEF_ACCESS_LOG_LEVEL = "WARN"

LOGGER = logging.getLogger(__package__)
LOGGER_FORMAT = "%(asctime)s [%(name)s] %(levelname)s %(message)s"


def main(args):
    """Main entry-point."""
    LOGGER.info("Starting application: version=%s", __version__)

    app = web.Application()
    app["client_proxy"] = args.client_proxy
    app["client_read_size"] = args.client_read_size
    app.add_routes((
        web.route("*", "/relay", relay_handler),
    ))
    app.cleanup_ctx.append(client_session_cleanup_ctx)

    LOGGER.info("Running HTTP server on %s:%d", args.host, args.port)
    web.run_app(app, host=args.host, port=args.port, print=None)


if __name__ == "__main__":
    PARSER = ArgumentParser(prog=__package__, description=__doc__,
                            formatter_class=ArgumentDefaultsHelpFormatter)

    HTTP_ARGS = PARSER.add_argument_group("main HTTP server arguments")
    HTTP_ARGS.add_argument("--host", metavar="HOSTNAME",
                           default=DEF_HOST,
                           help="HTTP server listening host")
    HTTP_ARGS.add_argument("--port", metavar="PORT",
                           default=DEF_PORT, type=int,
                           help="HTTP server listening port")

    RELAY_ARGS = PARSER.add_argument_group("relay HTTP client arguments")
    RELAY_ARGS.add_argument("--client-proxy", metavar="URL",
                            default=DEF_CLIENT_PROXY, type=URL,
                            help="HTTP proxy URL for the HTTP client")
    RELAY_ARGS.add_argument("--client-read-size", metavar="BYTES", type=int,
                            default=DEF_CLIENT_READ_SIZE,
                            help="data read size for the HTTP client")

    MONITORING_ARGS = PARSER.add_argument_group("application monitoring arguments")
    MONITORING_ARGS.add_argument("--log-level", metavar="LEVEL",
                                 default=DEF_LOG_LEVEL,
                                 help="application logging level")
    MONITORING_ARGS.add_argument("--access-log-level", metavar="LEVEL",
                                 default=DEF_ACCESS_LOG_LEVEL,
                                 help="HTTP server access logging level")

    ARGS = PARSER.parse_args()

    logging.basicConfig(format=LOGGER_FORMAT, level=ARGS.log_level)
    logging.getLogger("aiohttp.access").setLevel(ARGS.access_log_level)

    asyncio.set_event_loop_policy(uvloop.EventLoopPolicy())
    main(ARGS)
