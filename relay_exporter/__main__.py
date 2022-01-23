"""Simple Prometheus relay exporter application."""

import asyncio
import sys
from argparse import ArgumentParser, ArgumentDefaultsHelpFormatter
import uvloop
from aiohttp import web, ClientSession, ClientError


DEF_HOST = "0.0.0.0"
DEF_PORT = 9878


async def client_session_cleanup_ctx(app):
    """Cleanup context for an HTTP client session."""
    app["client_session"] = ClientSession(trust_env=True)
    yield
    await app["client_session"].close()


async def relay_handler(request):
    """Handle relay requests."""
    target = request.query.get("target")
    if not target:
        raise web.HTTPBadRequest(text="Target parameter is missing")
    client_session = request.config_dict["client_session"]
    try:
        async with client_session.get(target) as client_response:
            response = web.StreamResponse(
                status=client_response.status,
                headers=client_response.headers,
            )
            await response.prepare(request)
            async for line in client_response.content:
                await response.write(line)
    except ClientError as err:
        raise web.HTTPBadGateway(text=str(err)) from err


def main(args):
    """Application main entry-point."""
    app = web.Application()
    app.add_routes((
        web.get("/relay", relay_handler),
    ))
    app.cleanup_ctx.append(client_session_cleanup_ctx)
    web.run_app(app, host=args.host, port=args.port)


if __name__ == "__main__":
    PARSER = ArgumentParser(prog=__package__, description=__doc__,
                            formatter_class=ArgumentDefaultsHelpFormatter)

    PARSER.add_argument("-H", "--host", metavar="HOST", default=DEF_HOST,
                        help="TCP/IP host for the HTTP server")
    PARSER.add_argument("-p", "--port", metavar="PORT", default=DEF_PORT,
                        help="TCP/IP port for the HTTP server")

    ARGS = PARSER.parse_args()

    asyncio.set_event_loop_policy(uvloop.EventLoopPolicy())
    sys.exit(main(ARGS))
