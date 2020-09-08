import http from "../http-common";

class PlanetService {
  getAll(page, type, name) {
    //return http.get("/planet?page=" + encodeURIComponent(page) + "&type=" + encodeURIComponent(type));
    return http.get(`/planet/${page}/${type}/${name}`);
  }

  get(name) {
    return http.get(`/planet/${name}`);
  }

}

export default new PlanetService();