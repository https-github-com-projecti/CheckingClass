


class User {
  int id;
  String username;
  String password;

  User(int id, String username, String password) {
    this.id = id;
    this.username = username;
    this.password = password;
  }

  User.fromJson(Map json)
      : id = json['id'],
        username = json['username'],
        password = json['password'];

  Map toJson() {
    return {'id': id, 'username': username, 'password': password};
  }
}