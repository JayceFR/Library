from flask import Flask
from flask_restful import Api, Resource, reqparse
from flask_sqlalchemy import SQLAlchemy

app = Flask(__name__)

api = Api(app)
app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///database.db'
db = SQLAlchemy(app)

class UserModel(db.Model):
    id = db.Column(db.Integer, primary_key = True)
    email = db.Column(db.String(100), nullable = False)
    email = db.Column(db.String(100), nullable = False)

#with app.app_context():
#    db.create_all()

user_put_args = reqparse.RequestParser()
user_put_args.add_argument("email", type=str, help="Email Address")
user_put_args.add_argument("password", type=str, help="Password")


user = {}

class User(Resource):
    def get(self, id):
        return {"data": id}
    
    def put(self, id):
        args = user_put_args.parse_args()
        user[str(id)] = args
        return {id: user[str(id)]}

api.add_resource(User, "/user/<int:id>")


if __name__ == "__main__":
    app.run(debug=True)
