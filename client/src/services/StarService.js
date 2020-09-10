import http from "../http-common";

class StarService {
  getAll(page, classification, name) {
    return http.get(`/star/${page}/${classification}/${name}`);
  }

  get(name) {
    return http.get(`/star/${name}`);
  }

}

export default new StarService();