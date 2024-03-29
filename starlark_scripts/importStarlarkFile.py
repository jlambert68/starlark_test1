import sys
import types
import  pathlib
import typing


def importFileAs(
        modAsName: str,
        importedFilePath: typing.Union[str,  pathlib.Path],
) -> types.ModuleType:
    """ Import importedFilePath as modAsName, return imported module
by loading importedFilePath and registering modAsName in sys.modules.
importedFilePath can be any file and does not have to be a .py file. modAsName should be python valid.
Raises ImportError: If the file cannot be imported or any Exception: occuring during loading.

Refs:
Similar to: https://stackoverflow.com/questions/19009932/import-arbitrary-python-source-file-python-3-3
    but allows for other than .py files as well through importlib.machinery.SourceFileLoader.
    """
    import importlib.util
    import importlib.machinery

    # from_loader does not enforce .py but  importlib.util.spec_from_file_location() does.
    spec = importlib.util.spec_from_loader(
        modAsName,
        importlib.machinery.SourceFileLoader(modAsName, importedFilePath),
    )
    if spec is None:
        raise ImportError(f"Could not load spec for module '{modAsName}' at: {importedFilePath}")
    module = importlib.util.module_from_spec(spec)

    try:
        spec.loader.exec_module(module)
    except FileNotFoundError as e:
        raise ImportError(f"{e.strerror}: {importedFilePath}") from e

    sys.modules[modAsName] = module
    return module