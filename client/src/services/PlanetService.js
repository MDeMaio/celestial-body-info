import http from "../http-common";

class PlanetService {
  getAll(page, type, name) {
    return http.get(`/planet/${page}/${type}/${name}`);
  }

  get(name) {
    return http.get(`/planet/${name}`);
  }

  getPlanetTypes() {
    return http.get(`/planettype`);
  }

}

export default new PlanetService();