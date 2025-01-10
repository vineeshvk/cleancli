package templates

// 1: Directory or package name, 2: Group Name, 3: Provider Name, 4: Group Class Name
const FeatureDI = `
import 'package:%[1]s/feature/%[2]s/%[2]s_page_view_model.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

final %[3]sViewModuleProvider =
    ChangeNotifierProvider.autoDispose<%[4]sPageViewModel>(
  (ref) {
    return %[4]sPageViewModel();
  },
);
`

const FeatureRoutes = `
class LoginLocation extends BeamLocation<BeamState> {
  @override
  List<BeamPage> buildPages(BuildContext context, BeamState state) {
    return [
      CustomTransitionPage(
        key: const ValueKey("Login"),
        childWidget: const LoginPage(),
        pageTitle: "Login",
      ),
    ];
  }

  @override
  List<Pattern> get pathPatterns => ["/login"];
}
  `
