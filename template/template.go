package template

// 1: ApiMethod, 2: ResponseClass, 3: FunctionName, 4: ParamsWithTypeIfAny
const ApiServiceFunction = `
  @%s("%s")
  Future<HttpResponse<%s>> %s(%s);
`

// 1: PackageLocation
const ImportStatement = `import 'package:%s';`

// 1: ClassName
const DataSourceFileClass = `
import 'package:retrofit/retrofit.dart';

abstract class %sDS {
}
`

// 1: FolderOrGroupName, 2: ClassName
const DataSourceImplFileClass = `
import 'package:data/network/api_service.dart';
import 'package:data/source/%[1]s/%[1]s_data_source.dart';
import 'package:retrofit/retrofit.dart';

class %[2]sDSImpl extends %[2]sDS {
  final ApiService _apiService;

  %[2]sDSImpl(this._apiService);

}
`

// 1: ResponseClass, 2: FunctionName, 3: ParamsWithTypeIfAny
const DataSourceFileFunction = `
  Future<HttpResponse<%s>> %s(%s);
`

// 1: ResopnseClass, 2: FunctionName, 3: ParamsWithTypeIfAny, 4: ParamsVariable
const DataSourceImplFileFunction = `
  @override
  Future<HttpResponse<%s>> %[2]s(%s){
    return _apiService.%[2]s(%s);
  }
`
