"""Application signals module."""

from aiohttp import ClientSession


async def client_session(app):
    """Cleanup context for an HTTP client session."""
    app["client_session"] = ClientSession(trust_env=True)
    yield
    await app["client_session"].close()
