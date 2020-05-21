import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'edit.dart';

class MyHomePage extends StatefulWidget {
  MyHomePage({Key key, this.title}) : super(key: key);
  final String title;

  @override
  _MyHomePageState createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(widget.title),
      ),
      body: ListView(
        physics: BouncingScrollPhysics(),
        children: <Widget>[
          Row(
            children: <Widget>[
              Padding(
                  padding: EdgeInsets.only(left: 20, top: 20, bottom: 20),
                  child: Text('메모장',
                      style: TextStyle(fontSize: 36, color: Colors.blue))),
            ],
          ),
          ..._loadMemo(),
        ],
      ),
      floatingActionButton: FloatingActionButton.extended(
        onPressed: () {
          Navigator.push(
              context, CupertinoPageRoute(builder: (context) => EditPage()));
        },
        tooltip: '메모를 추가하려면 클릭하세요',
        label: Text('메모 추가'),
        icon: Icon(Icons.add),
      ), // This trailing comma makes auto-formatting nicer for build methods.
    );
  }

  List<Widget> _loadMemo() {
    List<Widget> memoList = [];

    memoList.add(Container(
      color: Colors.redAccent,
      height: 100,
    ));
    memoList.add(Container(
      color: Colors.orange,
      height: 100,
    ));
    memoList.add(Container(
      color: Colors.yellow,
      height: 100,
    ));
    memoList.add(Container(
      color: Colors.green,
      height: 100,
    ));
    memoList.add(Container(
      color: Colors.blue,
      height: 100,
    ));
    memoList.add(Container(
      color: Colors.indigo,
      height: 100,
    ));
    memoList.add(Container(
      color: Colors.purpleAccent,
      height: 100,
    ));

    return memoList;
  }
}
