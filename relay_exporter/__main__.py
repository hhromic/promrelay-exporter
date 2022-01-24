"""Simple Prometheus relay exporter entry-point module."""

import asyncio
import sys
from argparse import ArgumentParser, ArgumentDefaultsHelpFormatter
import uvloop
from aiohttp import web
from .handlers import relay_handler
from .signals import client_session


DEF_HOST = "0.0.0.0"
DEF_PORT = 9878
DEF_CLIENT_READ_SIZE = 2048


def main(args):
    """Application main entry-point."""
    app = web.Application()
    app["client_read_size"] = args.client_read_size
    app.add_routes((
        web.get("/relay", relay_handler),
    ))
    app.cleanup_ctx.append(client_session)
    web.run_app(app, host=args.host, port=args.port)


if __name__ == "__main__":
    PARSER = ArgumentParser(prog=__package__, description=__doc__,
                            formatter_class=ArgumentDefaultsHelpFormatter)

    PARSER.add_argument("-H", "--host", metavar="HOST", default=DEF_HOST,
                        help="TCP/IP host for the HTTP server")
    PARSER.add_argument("-p", "--port", metavar="PORT", default=DEF_PORT,
                        help="TCP/IP port for the HTTP server")
    PARSER.add_argument("-s", "--client-read-size", metavar="BYTES", default=DEF_CLIENT_READ_SIZE,
                        help="Data read size for the HTTP client")

    ARGS = PARSER.parse_args()

    asyncio.set_event_loop_policy(uvloop.EventLoopPolicy())
    sys.exit(main(ARGS))
