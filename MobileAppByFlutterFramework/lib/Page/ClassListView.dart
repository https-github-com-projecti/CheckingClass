import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:project1/Page/Subject.dart';
import 'package:project1/data/User.dart';
import '../main.dart';

class ClassListView {

  Card getStructuredGridCell(name,image,BuildContext context,User user,int tspassword){
    ScreenUtil.instance = ScreenUtil.getInstance()..init(context);
    
    return Card( 
      
      elevation: 5.0,
      shape: new RoundedRectangleBorder(borderRadius: new BorderRadius.circular(10.0)),
      color: Colors.limeAccent,
      child: InkWell(
        onTap: () {
          print("tapped");
          // Navigator.of(context).pushNamed("/" + subject);
          Navigator.push(context,MaterialPageRoute(builder: (context) => Subject(user: user,tspassword: tspassword,)));
        },
        child: Column(
          children: <Widget>[
            Padding(
              padding: const EdgeInsets.only(top: 80.0),
              child: Text(name,
                      style: TextStyle(
                        fontFamily: "Poppins-Bold",
                        fontSize: 20,
                        fontWeight: FontWeight.bold,),
                )
            ),
            // Text(x.toString()),
            // FlutterLogo(),
          ],
        ),
        ),
         
    );   
  }
  

 GridView build(BuildContext context, User user, {BuildContext itemBuilder}) {
   
    print("จำนวนวิชา \||/");
    print(subjectName.length);
    // print(subjectName[0].sName);r
    return GridView.count(
      // physics: ScrollPhysics(parent: ),
      // controller: controller,
      primary: true,
      padding: const EdgeInsets.all(2.0),
      crossAxisCount: 2,
      childAspectRatio: 0.85,
      mainAxisSpacing: 1.0,
      crossAxisSpacing: 0.5,
      
      children: <Widget>[
        // RefreshIndicator( 
        //   child:  getStructuredGridCell(subjectName[0].sName, "",context,user,subjectName[0].tspassword), 
        //   onRefresh: () {},
        //   color: Colors.red,

        // ),
        for (int i = 0; i < subjectName.length; i++) 
        getStructuredGridCell(subjectName[i].sName, "",context,user,subjectName[i].tspassword)
      ],
    );
  }


 
 
}

  


