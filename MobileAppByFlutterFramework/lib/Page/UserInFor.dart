import 'package:flutter/material.dart';
import 'package:project1/data/User.dart';
import 'package:project1/meterial/profile_tile.dart';

import '../main.dart';


class UserInFor extends StatelessWidget{
  final User user ;
  const UserInFor({Key key, this.user}) : super(key: key);

  @override 
  Widget build(BuildContext context){
    return Scaffold(
      appBar: AppBar( 
        elevation: 0,
        backgroundColor: Colors.white,
        iconTheme: new IconThemeData(color: Colors.black),
      ),
      backgroundColor: Colors.amber[numColor] ,
      body: Column(
        children: <Widget>[
          Padding( 
            padding: EdgeInsets.all(50.0),
            child: Center(
              child: Column( 
                children: <Widget>[ 
                  CircleAvatar(
                                backgroundImage: NetworkImage(
                                    "https://avatars0.githubusercontent.com/u/12619420?s=460&v=4"),
                                foregroundColor: Colors.black,
                                radius: 30.0,
                              ),
                  ProfileTile(
                    title: user.firstname +" " +user.lastname,
                    subtitle: user.idstudent,
                  ),

                  
                 
                ],
              ),
            ),
          ),
        ],
      ),
    );
  }
}