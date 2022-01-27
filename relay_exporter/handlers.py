"""Handlers module."""

import logging
from aiohttp import hdrs, web, ClientError, InvalidURL


LOGGER = logging.getLogger(__name__)


async def relay_handler(request):
    """Handler for relaying Prometheus scrape requests."""
    target = request.query.get("target")
    if not target:
        raise web.HTTPBadRequest(text="Target parameter is missing")

    client_session = request.config_dict["client_session"]
    client_proxy = request.config_dict["client_proxy"]
    client_read_size = request.config_dict["client_read_size"]
    try:
        client_headers = request.headers.copy()
        del client_headers[hdrs.HOST]

        client_args = {
            "headers": client_headers,
            "proxy": client_proxy,
        }

        LOGGER.debug("Relaying request to target: %s", target)
        async with client_session.get(target, **client_args) as client_response:
            request_response = web.StreamResponse(
                status=client_response.status,
                headers=client_response.headers,
            )
            await request_response.prepare(request)
            async for data in client_response.content.iter_chunked(client_read_size):
                await request_response.write(data)
            return request_response
    except InvalidURL as err:
        raise web.HTTPBadRequest(text=f"Invalid URL for target: {err}") from err
    except ClientError as err:
        raise web.HTTPBadGateway(text=f"Client error for target: {err}") from err
