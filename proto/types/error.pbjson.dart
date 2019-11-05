///
//  Generated code. Do not modify.
//  source: proto/types/error.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

const Error$json = const {
  '1': 'Error',
  '2': const [
    const {'1': 'code', '3': 1, '4': 1, '5': 5, '10': 'code'},
    const {'1': 'message', '3': 2, '4': 1, '5': 9, '10': 'message'},
    const {'1': 'details', '3': 4, '4': 1, '5': 11, '6': '.google.protobuf.Any', '10': 'details'},
    const {'1': 'validation', '3': 5, '4': 3, '5': 11, '6': '.types.Error.ValidationEntry', '10': 'validation'},
  ],
  '3': const [Error_ValidationEntry$json],
};

const Error_ValidationEntry$json = const {
  '1': 'ValidationEntry',
  '2': const [
    const {'1': 'key', '3': 1, '4': 1, '5': 9, '10': 'key'},
    const {'1': 'value', '3': 2, '4': 1, '5': 11, '6': '.types.ValidationError', '10': 'value'},
  ],
  '7': const {'7': true},
};

const ValidationError$json = const {
  '1': 'ValidationError',
  '2': const [
    const {'1': 'errors', '3': 1, '4': 3, '5': 9, '10': 'errors'},
  ],
};

