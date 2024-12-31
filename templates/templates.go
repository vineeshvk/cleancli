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

// 1: FolderOrGroupName, 2: GroupClassName, 3: DataDir
const DataSourceImplFileClass = `
import 'package:%[3]s/network/api_service.dart';
import 'package:%[3]s/source/%[1]s/%[1]s_data_source.dart';
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

// 1: GroupClassName, 2: DomainDir
const RepoFileClass = `
import 'package:dartz/dartz.dart';
import 'package:%[2]s/error/network_error.dart';

abstract class %[1]sRepository{
}
`

// 1: ResponseClass, 2: FunctionName, 3: ParamsWithTypeIfAny
const RepoFileFunction = `
  Future<Either<NetworkError, %s>> %s(%s);
`

// 1: FolderOrGroupName, 2: GroupClassName, 3: DataDir, 4: DomainDir
const RepoImplFileClass = `
import 'package:dartz/dartz.dart';
import 'package:%[3]s/network/safe_api_call.dart';
import 'package:%[3]s/source/%[1]s/%[1]s_data_source.dart';
import 'package:%[4]s/error/network_error.dart';
import 'package:%[4]s/repository/%[1]s_repository.dart';

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
// 6: params.toRequest(), 7: [UseCaseFileParamsToRequestFunction], 8: FolderOrGroupName, 9: DomainDir
const UseCaseFileClass = `
import 'package:dartz/dartz.dart';
import 'package:%[9]s/error/base_error.dart';
import 'package:%[9]s/usecase/base/base_usecase.dart';
import 'package:%[9]s/usecase/base/params.dart';
import 'package:%[9]s/repository/%[8]s_repository.dart';
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

// 1: GroupNameCamel, 2: GroupClassName
const DataSourceDI = `
final %[1]sDSProvider = Provider<%[2]sDS>(
  (ref) => %[2]sDSImpl(ref.read(apiServiceProvider))
);
`

// 1: GroupNameCamel, 2: GroupClassName
const RepoDI = `
final %[1]sRepoProvider = Provider<%[2]sRepository>(
  (ref) => %[2]sRepositoryImpl(ref.read(%[1]sDSProvider))
);
`

// 1: FunctionName, 2:FunctionPascal, 3: GroupNameCamel
const UseCaseDI = `
final %[1]sUseCaseProvider = Provider.autoDispose<%[2]sUseCase>(
  (ref) => %[2]sUseCase(ref.read(%[3]sRepoProvider)),
);
`

// 1: DataDir
const UseCaseDIInitImport = `
import 'package:%s/di/repository_modules.dart';
import 'package:riverpod/riverpod.dart';
`
