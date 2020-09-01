import http from "../http-common";

class PlanetService {
  getAll(page) {
    return http.get("/planet?page=" + page);
  }

  get(name) {
    return http.get(`/planet/${name}`);
  }

}

export default new PlanetService();