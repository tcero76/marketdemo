import { Meta, MetaRaw, Ts, Tses } from "@/types";
import { Servicios } from "@/types/servicios";

  const enrichMeta = (
    metaRaw: MetaRaw,
    tses: Tses[],
    servicios: Servicios[]
  ): Meta => {
    const normalize = (s?: string) => s ? s.slice(1).toLowerCase() : "";
    const normalizeNombre = (s: string) => s.replace(/\s+/g, "").toLowerCase();
    return {
      mentions: tses.filter(t =>
        metaRaw.mentions?.some(m => normalize(m) === normalizeNombre(t.nombre))
      ),
      hashtags: servicios.filter(s =>
        metaRaw.hashtags?.some(h => normalize(h) === normalizeNombre(s.nombre))
      ),
      urls: metaRaw.urls
    };
  }

const normalize = (s?: string) =>
  s ? s.slice(1).toLowerCase() : "";

const normalizeNombre = (s: string) =>
  s.replace(/\s+/g, "").toLowerCase();

function matchBy<T>(
  rawList: string[] | undefined,
  entities: T[],
  getKey: (item: T) => string
): T[] {
  if (!rawList?.length) return [];

  return entities.filter(e =>
    rawList.some(r => normalize(r) === normalizeNombre(getKey(e)))
  );
}

  export { enrichMeta, matchBy };