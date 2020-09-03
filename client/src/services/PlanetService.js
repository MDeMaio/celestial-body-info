import http from "../http-common";

class PlanetService {
  getAll(page, type) {
    return http.get("/planet?page=" + page + "&type=" + type);
  }

  get(name) {
    return http.get(`/planet/${name}`);
  }

}

export default new PlanetService();