package templates

// 1: Directory or package name, 2: Group Name, 3: Provider Name, 4: Group Class Name
const FeatureDI = `
import 'package:%[1]s/feature/%[2]s/%[2]s_page_view_model.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

final %[3]sModuleProvider =
    ChangeNotifierProvider.autoDispose<%[4]sPageViewModel>(
  (ref) {
    return %[4]sPageViewModel();
  },
);
`
// 1: Directory or package name, 2: Group Name
const FeatureRouteImports = `
import 'package:%[1]s/feature/%[2]s/%[2]s_page.dart';
`

// 1. Class name(PascalCase), 2. Name(with space in case of snake case), 3. route_name(in snake case)
const FeatureRoutes = `
class %[1]sLocation extends BeamLocation<BeamState> {
  @override
  List<BeamPage> buildPages(BuildContext context, BeamState state) {
    return [
      CustomTransitionPage(
        key: const ValueKey("%[2]s"),
        childWidget: const %[1]sPage(),
        pageTitle: "%[2]s",
      ),
    ];
  }

  @override
  List<Pattern> get pathPatterns => ["/%[3]s"];
}
  `

// 1. packageName, 2. group_name, 3. ClassName, 4. variableName
const FeaturePage = `
import 'package:%[1]s/base/base_page.dart';
import 'package:%[1]s/di/%[2]s_module.dart';
import 'package:%[1]s/feature/%[2]s/%[2]s_page_view.dart';
import 'package:%[1]s/feature/%[2]s/%[2]s_page_view_model.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

class %[3]sPage extends BasePage<%[3]sPageViewModel> {
  const %[3]sPage({super.key});

  static late ProviderBase<%[3]sPageViewModel> model;

  @override
  State<StatefulWidget> createState() => %[3]sPageState();
}

class %[3]sPageState extends BaseStatefulPage<
    %[3]sPageViewModel, %[3]sPage> {
  @override
  ProviderBase provideBase() {
    %[3]sPage.model = %[4]sModuleProvider;
    return %[3]sPage.model;
  }

  @override
  Widget buildView(BuildContext context, %[3]sPageViewModel model) {
    return %[3]sPageView(provideBase());
  }
}
`

// 1. packageName, 2. group_name, 3. ClassName
const FeaturePageView = `
import 'package:%[1]s/base/base_page.dart';
import 'package:%[1]s/feature/%[2]s/%[2]s_page_view_model.dart';
import 'package:flutter/material.dart';

class %[3]sPageView
    extends BasePageViewWidget<%[3]sPageViewModel> {
  const %[3]sPageView(super.provideBase, {super.key});

  @override
  Widget build(BuildContext context, model) {
    return Container();
  }
}
`

// 1. packageName, 2. ClassName
const FeaturePageViewModel = `
import 'package:%[1]s/base/base_page_view_model.dart';

class %[2]sPageViewModel extends BasePageViewModel
    with %[2]sPageViewModelStreams {
  %[2]sPageViewModel();
}

mixin %[2]sPageViewModelStreams {}
`
