from flask import Flask
from flask_restful import Api, Resource, reqparse, fields, marshal_with, abort
from flask_sqlalchemy import SQLAlchemy
from flask_cors import CORS

app = Flask(__name__)

#Solution to the cors error
CORS(app)
cors = CORS(app, resources={r"/api/*": {"origins": "*"}})

api = Api(app)
app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///database.db'
db = SQLAlchemy(app)


#All the DBS
class UserModel(db.Model):
    id = db.Column(db.Integer, primary_key = True)
    email = db.Column(db.String(100), nullable = False)
    password = db.Column(db.String(100), nullable = False)
    community_id = db.Column(db.Integer, nullable = True)

class CommunityModel(db.Model):
    id = db.Column(db.Integer, primary_key = True)
    name = db.Column(db.String(100), nullable = False)

create_db = False
if create_db:
    with app.app_context():
        db.create_all()


#Arguments Receiver
user_put_args = reqparse.RequestParser()
user_put_args.add_argument("email", type=str, help="Email Address")
user_put_args.add_argument("password", type=str, help="Password")
user_put_args.add_argument("community_id", type=str, help="Password")
#------------------------------------#
community_post_args = reqparse.RequestParser()
community_post_args.add_argument("name", type = str, help = "Name of community")

user = {}

user_resource_fields = {
    'id' : fields.Integer,
    'email' : fields.String,
    'password' : fields.String,
    'community_id': fields.Integer
}

community_resource_fields = {
    'id' : fields.Integer,
    'name': fields.String
}

class User(Resource):
    def get(self, id):
        results = UserModel.query.all()
        return_dict = {}
        for result in results:
            return_dict[str(result.id)] = result.email
        return return_dict
    
    @marshal_with(user_resource_fields)
    def put(self, id):
        args = user_put_args.parse_args()
        user = UserModel(id= id, email= args['email'], password = args['password'], community_id = args['community_id'])
        db.session.add(user)
        db.session.commit()
        return user

    #To update community of user
    #Used only after successful login 
    def post(self, id):
        args = user_put_args.parse_args()
        required_user = UserModel.query.filter_by(id = id).first()
        if not required_user:
            abort(404, message= "User does not exist")
        required_user.community_id = args['community_id']
        db.session.commit()

#Used to find an id for user
class UserIds(Resource):
    def get(self):
        users = UserModel.query.all()
        if len(users) > 0:
            user_id = users[len(users) - 1].id
        else:
            user_id = 0
        return {"id": user_id + 1 } 

class Community(Resource):
    def get(self, id):
        users = {}
    
    #Create a community
    @marshal_with(community_resource_fields)
    def post(self, id):
        args = community_post_args.parse_args()
        new_com = CommunityModel(id = id, name = args['name'])
        return new_com


api.add_resource(User, "/user/<int:id>")
api.add_resource(Community, "/community/<int:id>")
api.add_resource(UserIds, "/user_id")


if __name__ == "__main__":
    app.run(debug=True)
