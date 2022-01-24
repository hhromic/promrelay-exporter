"""Handlers module."""

from aiohttp import hdrs, web, ClientError


async def relay_handler(request):
    """Handler for relaying Prometheus scrape requests."""
    target = request.query.get("target")
    if not target:
        raise web.HTTPBadRequest(text="Target parameter is missing")
    client_session = request.config_dict["client_session"]
    client_read_size = request.config_dict["client_read_size"]
    try:
        client_headers = request.headers.copy()
        del client_headers[hdrs.HOST]
        async with client_session.get(target, headers=client_headers) as client_response:
            request_response = web.StreamResponse(
                status=client_response.status,
                headers=client_response.headers,
            )
            await request_response.prepare(request)
            async for data in client_response.content.iter_chunked(client_read_size):
                await request_response.write(data)
    except ClientError as err:
        raise web.HTTPBadGateway(text=f"Client error on target: {err}") from err
