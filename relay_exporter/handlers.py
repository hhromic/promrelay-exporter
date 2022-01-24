"""Handlers module."""

from aiohttp import web, ClientError


async def relay_handler(request):
    """Handler for relaying Prometheus scrape requests."""
    target = request.query.get("target")
    if not target:
        raise web.HTTPBadRequest(text="Target parameter is missing")
    client_session = request.config_dict["client_session"]
    client_read_size = request.config_dict["client_read_size"]
    try:
        async with client_session.get(target) as client_response:
            response = web.StreamResponse(
                status=client_response.status,
                headers=client_response.headers,
            )
            await response.prepare(request)
            async for data in client_response.content.iter_chunked(client_read_size):
                await response.write(data)
    except ClientError as err:
        raise web.HTTPBadGateway(text=f"{err}") from err
