import 'package:flutter/material.dart';
import 'package:project1/data/CreateUser.dart';

import 'package:project1/data/User.dart';
import 'package:project1/meterial/profile_tile.dart';

import '../main.dart';

TextEditingController phoneControler = new TextEditingController();
TextEditingController emailControler = new TextEditingController();
class UserInFor extends StatelessWidget{
  final User user ;
  
  const UserInFor({Key key, this.user}) : super(key: key);
  

  @override 
  Widget build(BuildContext context){
    return Scaffold(
      appBar: AppBar( 
        elevation: 1,
        backgroundColor: Colors.white,
        iconTheme: new IconThemeData(color: Colors.black),
      ),
      backgroundColor: Colors.amber[numColor1] ,
      body: Column(
        children: <Widget>[
          Padding( 
            padding: EdgeInsets.only(top: 50.0),
            child: Center(
              child: Column( 
                children: <Widget>[ 
                  CircleAvatar(
                                backgroundImage: NetworkImage(
                                    "https://images-wixmp-ed30a86b8c4ca887773594c2.wixmp.com/f/89566ebc-ae66-4173-858c-bdae4ee22eb2/dap48o3-b6df2aed-517b-4feb-96b5-33f0d2677c00.png?token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ1cm46YXBwOjdlMGQxODg5ODIyNjQzNzNhNWYwZDQxNWVhMGQyNmUwIiwiaXNzIjoidXJuOmFwcDo3ZTBkMTg4OTgyMjY0MzczYTVmMGQ0MTVlYTBkMjZlMCIsIm9iaiI6W1t7InBhdGgiOiJcL2ZcLzg5NTY2ZWJjLWFlNjYtNDE3My04NThjLWJkYWU0ZWUyMmViMlwvZGFwNDhvMy1iNmRmMmFlZC01MTdiLTRmZWItOTZiNS0zM2YwZDI2NzdjMDAucG5nIn1dXSwiYXVkIjpbInVybjpzZXJ2aWNlOmZpbGUuZG93bmxvYWQiXX0.5-IneqmHJLinXfaDtuwCdTtA6MNeBzwr-LlYQAx-V9I"),
                                foregroundColor: Colors.black,
                                radius: 30.0,
                              ),
                  ProfileTile(
                    title: user.firstname +" " +user.lastname,
                    subtitle: user.idstudent,
                  ),
                  //  Divider(
                  //   color: Colors.black,
                  //   height: 8.0,
                  // ),
                  Padding(
                    padding: const EdgeInsets.only(top: 50.0,right: 8.0,left: 8.0),
                    child: Row(
                      children: <Widget>[
                        Icon(Icons.info),
                        Text(" Student Information",
                          style: TextStyle(fontSize: 25.0,fontWeight: FontWeight.bold),
                        ),
                      ],
                    ),
                  ),
                 
                   Divider(
                    color: Colors.black,
                    height: 8.0,
                  ),
                   Padding(
                     padding: const EdgeInsets.only(top:20.0,right: 8.0,left: 8.0,bottom: 20.0),
                     child: Row(
                       children: <Widget>[
                         Icon(Icons.call),
                         Text(" Phone: ",
                            style: TextStyle(fontSize: 20.0,fontWeight: FontWeight.bold),
                         ),
                         Text(user.phone,
                            style: TextStyle(fontSize: 20.0,),
                         ),
                         IconButton(
                           icon: Icon(Icons.edit,),
                           color: Colors.greenAccent,
                           onPressed: (){
                             showDialog(
                              context: context,
                              builder: (context) => new AlertDialog(
                                shape: new RoundedRectangleBorder(borderRadius: new BorderRadius.circular(10.0)),
                                title: new Text("Edit Phone"),
                                content: Container(
                                    height: 150.0, 
                                    child: Column(  
                                      children: <Widget>[ 
                                        // Text("Coming Soon..."),
                                        TextField(
                                          controller: phoneControler,
                                          maxLength: 10,
                                          keyboardType: TextInputType.number,
                                          decoration: InputDecoration(
                                            border: OutlineInputBorder(borderRadius: const BorderRadius.all(Radius.circular(20.0)), borderSide: BorderSide(color: Colors.amber,),),
                                            focusedBorder:OutlineInputBorder(borderSide: const BorderSide(color: Colors.teal, width: 2.0),borderRadius: BorderRadius.circular(20.0),),
                                            hintText: user.phone,
                                            hintStyle: TextStyle(color: Colors.grey, fontSize: 12.0)
                                          ),
                                        ),
                                        FlatButton( 
                                          child: Text("Edit"),
                                          color: Colors.greenAccent,
                                          onPressed: () async {
                                              CreateUser edit = new CreateUser(firstname: user.firstname, lastname: user.lastname,
                                                              idstudent: user.idstudent, phone: phoneControler.text, username: user.username, password: user.password,
                                                              email: user.email,);
                                              CreateUser p = await editPhoneAndEmail(putPhoneURL,body: edit.toMap());
                                              if (p != null){
                                                print("--------------NULL-------------");
                                              }
                                          },
                                        )
                                      ],
                                    ),
                                  ),                   
                                )
                              );
                           },
                          )
                       ],
                     ),
                   ),
                    Divider(
                    color: Colors.black,
                    height: 8.0,
                  ),
              
                   Padding(
                     padding: const EdgeInsets.only(top:20.0,right: 8.0,left: 8.0,bottom: 20.0),
                     child: Row(
                       children: <Widget>[
                         Icon(Icons.email),
                         Text(" E-mail: ",
                            style: TextStyle(fontSize: 20.0,fontWeight: FontWeight.bold),
                         ),
                          Text(user.email,
                            style: TextStyle(fontSize: 20.0,),
                         ),
                         IconButton(
                           icon: Icon(Icons.edit,),
                           color: Colors.greenAccent,
                           onPressed: (){
                             showDialog(
                              context: context,
                              builder: (context) => new AlertDialog(
                                shape: new RoundedRectangleBorder(borderRadius: new BorderRadius.circular(10.0)),
                                title: new Text("Edit E-mail"),
                                content: Container(
                                    height: 150.0, 
                                    child: Column(  
                                      children: <Widget>[ 
                                        // Text("Coming Soon...")
                                        TextField(
                                          controller: emailControler,
                                          keyboardType: TextInputType.emailAddress,
                                          decoration: InputDecoration(
                                            border: OutlineInputBorder(borderRadius: const BorderRadius.all(Radius.circular(20.0)), borderSide: BorderSide(color: Colors.amber,),),
                                            focusedBorder:OutlineInputBorder(borderSide: const BorderSide(color: Colors.teal, width: 2.0),borderRadius: BorderRadius.circular(20.0),),
                                            hintText: user.email,
                                            hintStyle: TextStyle(color: Colors.grey, fontSize: 12.0)
                                          ),
                                        ),
                                        Padding(
                                          padding: const EdgeInsets.only(top:15.0),
                                          child: FlatButton( 
                                            child: Text("Edit"),
                                            color: Colors.greenAccent,
                                            // shape: ,
                                            onPressed: () async {
                                              CreateUser edit = new CreateUser(firstname: user.firstname, lastname: user.lastname,
                                                              idstudent: user.idstudent, phone: user.phone, username: user.username, password: user.password,
                                                              email: emailControler.text,);
                                              CreateUser p = await editPhoneAndEmail(putemailURL,body: edit.toMap());
                                              if (p != null){
                                                print("--------------NULL-------------");
                                              }
                                            },
                                          ),
                                        )
                                      ],
                                    ),
                                  ),                   
                                )
                              );
                           },
                          )
                       ],
                     ),
                   ),
                      Divider(
                    color: Colors.black,
                    height: 8.0,
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