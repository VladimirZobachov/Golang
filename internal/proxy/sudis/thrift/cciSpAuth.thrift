namespace cpp sudis
namespace java ru.atc.mvd.sudis.thrift.cci

enum TCciSpAuthVersion {
    // версия 1
    V1 = 1
}

enum TCciSpAuthOperationResult {
    // Операция успешно выполнена
    SUCCESS = 1,
    // Внутренняя ошибка сервера
    ERROR_INTERNAL = 2,
    // Повторная отправка запроса запрещена
    REQUEST_REPLAY_RESTRICTED = 3,
    // Запрос устарел (возможно несоответствие времени на отправителе и получателе)
    // Допустимая разница во времени между отправителем получателем - 3 минуты (180000 миллисекунд)
    REQUEST_EXPIRED = 4,
    // Неправильный идентификатор запроса (null или длина менее 4 байт или длина более 32 байт)
    REQUEST_NONCE_NOT_VALID = 5,
    // Недостаточно полномочий сервиса для вызова данного метода
    SP_PERMISSION_DENIED = 6,
    // Учетная запись сервиса заблокирована в СУДИС
    SP_ACCOUNT_BLOCKED = 7,
    // Ошибка аутентификации сервиса ИСОД
    SP_NOT_AUTHENTICATED = 8
}

struct TokenDataPerm{
    //Код сервиса
    1: required string spCode;
    //Название сервиса
    2: required string spName;
}

struct TCciSpAuthCreateTokenArgs {
    // Версия структуры
    10: optional TCciSpAuthVersion version = TCciSpAuthVersion.V1,
    // время создания запроса
    // для предотвращения пересылки старых запросов
    // разница во времени, измеренная в миллисекундах,
    // между заданным временем и полночью 01.01.1970 в часовом поясе UTC.
    20: optional i64 requestMillis = 0,
    // случайная байтовая последовательность длиной не менее 4 байт
    // для идентификации запроса
    30: optional binary requestNonce,
    //код Сервиса ИСОД, запрашивающий доступ
    40: optional string spCode,
    // код Сервиса ИСОД, к которому будет осуществляться доступ
    50: optional string targetSpCode
}

struct TCciSpAuthTokenDataArgs {
    // Версия структуры
    10: optional TCciSpAuthVersion version = TCciSpAuthVersion.V1,
    // время создания запроса
    // для предотвращения пересылки старых запросов
    // разница во времени, измеренная в миллисекундах,
    // между заданным временем и полночью 01.01.1970 в часовом поясе UTC.
    20: optional i64 requestMillis = 0,
    // случайная байтовая последовательность длиной не менее 4 байт
    // для идентификации запроса
    30: optional binary requestNonce,
    // Идентификатор токена
    40: optional binary tokenId;
}

struct TCciSpAuthTokenData {
    // Версия структуры
    10: optional TCciSpAuthVersion version = TCciSpAuthVersion.V1,
    // Идентификатор токена
    20: optional binary tokenId;
    // время устаревания токена
    // разница во времени, измеренная в миллисекундах,
    // между заданным временем и полночью 01.01.1970 в часовом поясе UTC.
    30: optional i64 expireMillis = 0,
    // код Сервиса ИСОД, для которого создан токен
    40: optional string spCode,
    // код Сервиса ИСОД, для доступа к которому создан токен
    50: optional string targetSpCode,
    // полномочия доступа к целевому сервису ИСОД
    60: optional list<TokenDataPerm> permissions
}

struct TCciSpAuthTokenResult {
    10: optional TCciSpAuthVersion version = TCciSpAuthVersion.V1,
    // время создания ответа
    // для предотвращения пересылки старых ответов
    // разница во времени, измеренная в миллисекундах,
    // между заданным временем и полночью 01.01.1970 в часовом поясе UTC.
    20: optional i64 responseMillis = 0,
    // случайная байтовая последовательность длиной не менее 4 байт и не более 32 байт
    // для идентификации ответа
    30: optional binary responseNonce,
    // результат операции
    40: optional string result = "SUCCESS",
    // дополнительное сообщение о результате операции
    50: optional string resultMessage = "",
    // Данные токена
    60: optional TCciSpAuthTokenData tokenData;
}

service TCciSpAuth {
    // запрос на создание токена
    TCciSpAuthTokenResult createToken(
        1: TCciSpAuthCreateTokenArgs arguments
    ),
    // запрос на получение данных токена
    TCciSpAuthTokenResult tokenData(
        1: TCciSpAuthTokenDataArgs arguments
    )
}
