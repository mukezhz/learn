import httpx
from mcp.server.fastmcp import FastMCP

OSM_API_KEY = "open-weather-map-key"
OWM_URL = "https://api.openweathermap.org/data/2.5/weather"

mcp = FastMCP(
    name="wesionaryTEAM weather info",
    dependencies=["httpx"],
)


@mcp.tool(name="get_weather_app")
def get_weather_by_city_name(city: str) -> dict[str, object]:
    params = {
        "q": city,
        "appid": OSM_API_KEY,
        "units": "metric",
    }

    response = httpx.get(OWM_URL, params=params)

    if response.status_code == 200:
        return response.json()
    else:
        return {
            "error": response.status_code,
            "message": response.text,
        }


if __name__ == "__main__":
    mcp.run(
        "stdio",
    )
