import http from "../http-common";

class StarService {
  getAll(page, type) {
    return http.get("/star?page=" + encodeURIComponent(page) + "&type=" + encodeURIComponent(type));
  }

  get(name) {
    return http.get(`/star/${name}`);
  }

}

export default new StarService();