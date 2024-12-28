package templates

// 1: ApiMethod, 2: ResponseClass, 3: FunctionName, 4: ParamsWithTypeIfAny
const ApiServiceFunction = `
  @%s("%s")
  Future<HttpResponse<%s>> %s(%s);
`

// 1: PackageLocation
const ImportStatement = `import 'package:%s';`

// 1: GroupClassName
const DataSourceFileClass = `
import 'package:retrofit/retrofit.dart';

abstract class %sDS {
}
`

// 1: FolderOrGroupName, 2: GroupClassName
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
  Future<HttpResponse<%s>> %[2]s(%[3]s){
    return _apiService.%[2]s(%[4]s);
  }
`

// 1: GroupClassName
const RepoFileClass = `
import 'package:dartz/dartz.dart';
import 'package:domain/error/network_error.dart';

abstract class %sRepository{
}
`

// 1: ResponseClass, 2: FunctionName, 3: ParamsWithTypeIfAny
const RepoFileFunction = `
  Future<Either<NetworkError, %s>> %s(%s);
`

// 1: FolderOrGroupName, 2: GroupClassName
const RepoImplFileClass = `
import 'package:dartz/dartz.dart';
import 'package:data/network/safe_api_call.dart';
import 'package:data/source/%[1]s/%[1]s_data_source.dart';
import 'package:domain/error/network_error.dart';
import 'package:domain/repository/%[1]s_repository.dart';

class %[2]sRepositoryImpl implements %[2]sRepository{
  final %[2]sDS _remoteDS;

  %[2]sRepositoryImpl(this._remoteDS);
}
`

// 1: ResponseClass, 2: FunctionName, 3:ParamsWithTypeIfAny, 4: ParamsVariable
const RepoImplFileFunction = `
  @override
  Future<Either<NetworkError, %[1]s>> %[2]s(%[3]s) async {
    final result = await safeApiCall(_remoteDS.%[2]s(%[4]s));

    return result!.fold((l) => Left(l), (r) => Right(r.data));
  }
`

// 1: ImportList, 2:FunctionPascal, 3: ResponseClass, 4: GroupClassName, 5: FunctionName,
// 6: params.toRequest(), 7: [UseCaseFileParamsToRequestFunction] 8: FolderOrGroupName
const UseCaseFileClass = `
import 'package:dartz/dartz.dart';
import 'package:domain/error/base_error.dart';
import 'package:domain/usecase/base/base_usecase.dart';
import 'package:domain/usecase/base/params.dart';
import 'package:domain/repository/%[8]s_repository.dart';
%[1]s

class %[2]sUseCase extends BaseUseCase<BaseError, %[2]sUseCaseParams,  %[3]s> {
  final %[4]sRepository _repo;

  %[2]sUseCase(this._repo);

  @override
  Future<Either<BaseError, %[3]s>> execute(
      {required %[2]sUseCaseParams params}) async {
    return await _repo.%[5]s(%[6]s);
  }
}

class %[2]sUseCaseParams extends Params {
  %[2]sUseCaseParams();
  %[7]s
}
`

// 1: RequestClass
const UseCaseFileParamsToRequestFunction = `
  %[1]s toRequest(){
    return %[1]s();
  }
`
