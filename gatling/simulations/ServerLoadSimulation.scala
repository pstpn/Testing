import io.gatling.core.Predef._
import io.gatling.http.Predef._
import scala.concurrent.duration._

class ServerLoadSimulation extends Simulation {

  val httpProtocolEcho = http.baseUrl("http://collector:8081")

  val scnEcho = scenario("Echo Server Load Test")
    .exec(http("Echo Metrics").get("/ping"))

  setUp(
    scnEcho.inject(
      rampConcurrentUsers(0).to(100).during(10.seconds),
      constantConcurrentUsers(100).during(10.seconds)
    ).protocols(httpProtocolEcho)
  ).maxDuration(20.seconds)
}