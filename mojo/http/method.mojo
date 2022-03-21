
@case_style("screaming_snake")
enum Method {
    unsepecified @0

    /// The HTTP GET method requests a representation of the specified resource.
    ///
    /// Requests using GET should only be used to request data (they shouldn't include data).
    get @1

    post @2

    put @3

    patch @4

    delete @5

    options @6

    head @7

    trace @8

    connect @9
}
