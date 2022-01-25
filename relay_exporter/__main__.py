"""Simple Prometheus relay exporter entry-point module."""

import asyncio
import logging
import sys
from argparse import ArgumentParser, ArgumentDefaultsHelpFormatter
import uvloop
from aiohttp import web
from .handlers import relay_handler
from .signals import client_session
from .version import __version__


DEF_HOST = "0.0.0.0"
DEF_PORT = 9878
DEF_CLIENT_READ_SIZE = 2048
DEF_LOG_LEVEL = "INFO"


LOGGER = logging.getLogger(__name__)
LOGGER_FORMAT = "%(asctime)s [%(name)s] %(levelname)s %(message)s"


def main(args):
    """Application main entry-point."""
    LOGGER.info("Starting application. Version: %s", __version__)

    app = web.Application()
    app["client_read_size"] = args.client_read_size
    app.add_routes((
        web.get("/relay", relay_handler),
    ))
    app.cleanup_ctx.append(client_session)

    LOGGER.info("Running HTTP server on %s:%d", args.host, args.port)
    web.run_app(app, host=args.host, port=args.port, print=None)


if __name__ == "__main__":
    PARSER = ArgumentParser(prog=__package__, description=__doc__,
                            formatter_class=ArgumentDefaultsHelpFormatter)

    PARSER.add_argument("-H", "--host", metavar="HOST", default=DEF_HOST,
                        help="TCP/IP host for the HTTP server")
    PARSER.add_argument("-p", "--port", metavar="PORT", default=DEF_PORT,
                        help="TCP/IP port for the HTTP server")
    PARSER.add_argument("-s", "--client-read-size", metavar="BYTES", default=DEF_CLIENT_READ_SIZE,
                        help="Data read size for the HTTP client")
    PARSER.add_argument("-l", "--log-level", metavar="LEVEL", default=DEF_LOG_LEVEL,
                        help="Application logging level")

    ARGS = PARSER.parse_args()

    logging.basicConfig(format=LOGGER_FORMAT, level=ARGS.log_level)
    asyncio.set_event_loop_policy(uvloop.EventLoopPolicy())
    sys.exit(main(ARGS))
