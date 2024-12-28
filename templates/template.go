package templates

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

// 1: ResponseClass, 2: FunctionName, 3: ParamsWithTypeIfAny, 4: ParamsVariable
const DataSourceImplFileFunction = `
  @override
  Future<HttpResponse<%s>> %[2]s(%s){
    return _apiService.%[2]s(%s);
  }
`

// 1: ClassName
const RepoFileClass = `
import 'package:dartz/dartz.dart';
import 'package:domain/error/network_error.dart';

abstract class %sRepository{
}
`

// 1: ResponseClass, 2: FunctionName, 3: ParamsWithTypeIfAny
const RepoFileFunction = `
  Future<Either<NetworkError, %s> %s(%s);
`

// 1: FolderOrGroupName, 2: ClassName
const RepoImplFileClass = `
import 'package:dartz/dartz.dart';
import 'package:data/network/safe_api_call.dart';
import 'package:domain/error/network_error.dart';
import 'package:domain/repository/%s_repository.dart';
class %[2]sRepositoryImpl implements %[2]sRepository{
  final %[2]sDS _remoteDS;

  %[2]sRepositoryImpl(this._remoteDS);
}
`

// 1: ResponseClass, 2: FunctionName, 3:ParamsWithTypeIfAny, 4: ParamsVariable
const RepoImplFileFunction = `
  @override
  Future<Either<NetworkError, %s>> %[2]s(%s) async {
    final result = await safeApiCall(_remoteDS.%[2]s(%s));

    return result!.fold((l) => Left(l), (r) => Right(r.data));
  }
`
