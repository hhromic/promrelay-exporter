"""Application signals module."""

import logging
from aiohttp import ClientSession


LOGGER = logging.getLogger(__name__)


async def client_session(app):
    """Cleanup context for an HTTP client session."""
    LOGGER.info("Creating application HTTP client session")
    app["client_session"] = ClientSession(auto_decompress=False)
    yield
    await app["client_session"].close()
    LOGGER.info("Closed application HTTP client session")
