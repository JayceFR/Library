from flask import Flask
from flask_restful import Api, Resource, reqparse

app = Flask(__name__)

api = Api(app)

user_put_args = reqparse.RequestParser()
user_put_args.add_argument("email", type=str, help="Email Address")
user_put_args.add_argument("password", type=str, help="Password")


class User(Resource):
    def get(self, id):
        return {"data": id}
    
    def put(self, id):
        args = user_put_args.parse_args()
        return {id: args}

api.add_resource(User, "/user/<int:id>")


if __name__ == "__main__":
    app.run(debug=True)
