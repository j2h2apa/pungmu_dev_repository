import 'package:flutter/material.dart';
import 'package:isolate_flutter/screens/performance_screen.dart';
import 'package:isolate_flutter/screens/infinite_process_screen.dart';
import 'package:isolate_flutter/screens/data_transfer_screen.dart';

class HomeScreen extends StatelessWidget {
  static const id = 'home_screen';
  @override
  Widget build(BuildContext context) {
    return DefaultTabController(
      length: 3,
      child: Scaffold(
        appBar: AppBar(
          bottom: TabBar(
            tabs: [
              Tab(
                icon: Icon(Icons.flash_on),
                text: 'Performance',
              ),
              Tab(
                icon: Icon(Icons.sync),
                text: 'Infinite Process',
              ),
              Tab(
                icon: Icon(Icons.storage),
                text: 'Data Transfer',
              ),
            ],
          ),
          title: Text('Isolate Tabs'),
        ),
        body: TabBarView(
          children: [
            PerformanceScreen(),
            InifiniteProcessScreen(),
            DataTranferScreen()
          ],
        ),
      ),
    );
  }
}
