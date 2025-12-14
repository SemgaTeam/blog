import { DataProvider } from "react-admin";
import { config } from "./config.ts";

const dataProvider: DataProvider = {
  getList: async (resource, params) => {
    let request = `${config.apiUrl}/${resource}?`;

    if (params.pagination.page && params.pagination.perPage) {
      request += `pagination=${JSON.stringify(params.pagination)}&`;
    }

    if (params.sort.field) {
      const sorting = {
        sortField: params.sort.field,
        sortOrder: params.sort.order,
      };

      request += `sorting=${JSON.stringify(sorting)}&`;
    }

    const response = await fetch(request);
    const data = await response.json();

    return {
      data: data.data || [],
      total: data.total,
    };
  },

  getOne: async (resource, params) => {
    const request = `${config.apiUrl}/${resource}/${params.id}`;

    const response = await fetch(request);
    const data = await response.json();

    return {
      data: data,
    };
  },

  getMany: async (resource, params) => {
    let ids = "";
    for (const id of params.ids) {
      ids += `ids=${id}&`;
    }

    const request = `${config.apiUrl}/${resource}?${ids}`;

    const response = await fetch(request);
    const data = await response.json();

    return {
      data: data.data,
    };
  },

  getManyReference: async (resource, params) => {
    let request = `${config.apiUrl}/${resource}?${params.target}=${params.id}&`;

    if (params.pagination.page && params.pagination.perPage) {
      request += `pagination=${JSON.stringify(params.pagination)}&`;
    }

    if (params.sort.field && params.sort.order) {
      request += `sorting=${JSON.stringify(sorting)}&`;
    }

    const response = await fetch(request);
    const data = await response.json();

    return {
      data: data,
      total: data.length,
    };
  },

  create: async (resource, params) => {
    const request = `${config.apiUrl}/${resource}`;

    const response = await fetch(request, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(params.data),
    });

    const data = await response.json();
    return { data };
  },

  update: async (resource, params) => {
    const request = `${config.apiUrl}/${resource}/${params.id}`;

    const response = await fetch(request, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(params.data),
    });

    const data = await response.json();
    return { data: data };
  },

  updateMany: async (resource, params) => {
    const responses = await Promise.all(
      params.ids.map((id) => {
        fetch(`${config.apiUrl}/${resource}/${id}`, {
          method: "PUT",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(params.data),
        }).then((res) => res.json());
      }),
    );
    return { data: responses.map((r) => r.id) };
  },

  delete: async (resource, params) => {
    await fetch(`${config.apiUrl}/${resource}/${params.id}`, {
      method: "DELETE",
    });

    return { data: { id: params.id } };
  },

  deleteMany: async (resource, params) => {
    await Promise.all(
      params.ids.map((id) => {
        fetch(`${config.apiUrl}/${resource}/${id}`, { method: "DELETE" });
      }),
    );
    return { data: params.ids };
  },
};

export default dataProvider;
