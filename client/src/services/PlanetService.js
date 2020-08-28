import http from "../http-common";

class PlanetService {
  getAll() {
    return http.get("/planet");
  }

  get(id) {
    return http.get(`/planet/${id}`);
  }

}

export default new PlanetService();