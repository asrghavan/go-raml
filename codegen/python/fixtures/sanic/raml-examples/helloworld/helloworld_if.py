from sanic import Blueprint
from sanic.views import HTTPMethodView
from sanic.response import text
import helloworld_api


helloworld_if = Blueprint('helloworld_if')


class helloworldView(HTTPMethodView):

    async def get(self, request):

        return await helloworld_api.helloworld_get(request)


helloworld_if.add_route(helloworldView.as_view(), '/helloworld')
